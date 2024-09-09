import { NextResponse } from "next/server";
import type { NextRequest } from "next/server";
import { cookies } from "next/headers";
import { revalidatePath } from "next/cache";

export function middleware(request: NextRequest) {
  const cookie = cookies().get("session")?.value;
  const role = cookies().get("role")?.value;

  if (cookie === undefined && request.nextUrl.pathname !== "/login") {
    console.log("Redirect to login");
    return NextResponse.redirect(new URL("/login", request.url));
  } else if (cookie !== undefined && request.nextUrl.pathname === "/login") {
    console.log("Redirect to restaurants");
    return NextResponse.redirect(new URL("/restaurants", request.url));
  } else if (request.nextUrl.pathname.includes("admin") && role !== "admin") {
    console.log("Not an admin");
    return NextResponse.redirect(new URL("/restaurants", request.url));
  }

  return NextResponse.next();
}

export const config = {
  matcher: ["/admin/:path*", "/restaurants/:path*", "/login"],
};
