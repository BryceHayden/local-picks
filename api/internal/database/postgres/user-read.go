package postgres

import (
	_ "github.com/lib/pq"

	"github.com/brycehayden/resume/internal/models"
)

func (psql PSQL) FetchUser(email string) (*models.User, error) {
	user := &models.User{};

    err := psql.DB.Get(
		&user, 
		`select * from member where email = $1;`, 
		email,
	)
	
	if err != nil {
		return nil, err
	}
	
	return user, nil
}


func (psql PSQL) FetchUserRole(email string) (*models.UserRole, error) {
	user := &models.UserRole{};

    err := psql.DB.Get(
		&user, 
		`
			select m.member_id, m.first_name, m.middle_name, m.surname, m.email, 
				CASE 
					WHEN mr.role_id is not distinct from null THEN 'base'
					ELSE r.name
				END as role
				from member m
				left join member_role mr on mr.member_id = m.member_id
				left join roles r on r.role_id = mr.role_id
			where m.email = $1
			group by m.member_id, m.first_name, m.middle_name, m.surname, m.email, mr.role_id, r.name;
		`, 
		email)
	
	if err != nil {
		return nil, err
	}
	
	return user, nil
}

func (psql PSQL) FetchUserRolePasshash(email string) (*models.UserRolePasshash, error) {
	user := &models.UserRolePasshash{};
	err := psql.DB.QueryRowx(
		`
			select m.member_id, m.first_name, m.middle_name, m.surname, m.email, m.joined_date, m.username, p.passhash, 
				CASE 
					WHEN mr.role_id is not distinct from null THEN 'base'
					ELSE r.name
				END as role
				from member m
				left join member_role mr on mr.member_id = m.member_id
				left join roles r on r.role_id = mr.role_id
				left join passhash p on p.member_id = m.member_id
			where m.email=$1
			group by m.member_id, m.username, m.first_name, m.middle_name, m.surname, m.email, m.joined_date, mr.role_id, r.name, p.passhash
			limit 1;
		`, 
		email,
	).StructScan(user)
	
	if err != nil {
		return nil, err
	}
	
	return user, nil
}
