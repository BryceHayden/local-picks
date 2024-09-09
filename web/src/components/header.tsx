import Link from "next/link";
import { cookies } from "next/headers";
import { logoutUser } from "@/server/login";

export default function Header() {
  const session =
    cookies().get("session")?.value !== undefined &&
    cookies().get("session")?.value !== "";
  const admin =
    cookies().get("role")?.value && cookies().get("role")?.value === "admin";

  return (
    <header className="flex items-center w-full h-24 px-16 bg-stone-300 border-b-2 border-slate-400">
      <nav className="flex justify-between items-center w-full">
        <Link
          href="/"
          className="text-4xl font-bold tracking-wide"
          style={{ fontFamily: "Montserrat" }}
        >
          Local Picks
        </Link>

        <div className="text-2xl flex items-center h-12 gap-4 print:hidden">
          {session ? (
            <>
              {admin ? <Link href="/admin">Admin</Link> : null}
              <Link href="/restaurants">Restaurants</Link>
              <form action={logoutUser}>
                <button type="submit">Logout</button>
              </form>
            </>
          ) : (
            <Link href="/login">Login</Link>
          )}
        </div>
      </nav>
    </header>
  );
}
