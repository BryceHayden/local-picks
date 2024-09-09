import { loginUser } from "@/server/login";

export default function Login() {
  return (
    <main className="flex flex-col gap-8 row-start-2 items-center sm:items-start">
      <form
        action={loginUser}
        className="flex flex-col items-center justify-center"
      >
        <div className="flex items-center justify-evenly gap-12 pb-6">
          <label>
            <div>Email</div>
            <input
              className="bg-gray-50 border border-gray-300 text-black text-sm rounded-lg focus:ring-blue-500 focus:border-blue-500 block w-full p-2.5 dark:bg-gray-300 dark:border-gray-500 dark:placeholder-blue-500 dark:text-black dark:focus:ring-blue-500 dark:focus:border-blue-500"
              name="email"
              placeholder="Email"
              type="email"
            />
          </label>
          <label>
            <div>Password</div>
            <input
              className="bg-gray-50 border border-gray-300 text-black text-sm rounded-lg focus:ring-blue-500 focus:border-blue-500 block w-full p-2.5 dark:bg-gray-300 dark:border-gray-500 dark:placeholder-blue-500 dark:text-black dark:focus:ring-blue-500 dark:focus:border-blue-500"
              name="password"
              placeholder="Password"
              type="password"
            />
          </label>
        </div>
        <button type="submit" className="bg-blue-300 rounded-lg px-6 py-2">
          Let's Go!
        </button>
      </form>
    </main>
  );
}
