"use server";
import { cookies } from "next/headers";
import { redirect } from "next/navigation";

export const loginUser = async (formData: FormData) => {
  try {
    const loginResponse = await fetch("http://api:9000/v1/login", {
      body: JSON.stringify({
        email: formData.get("email"),
        password: formData.get("password"),
      }),
      method: "POST",
    });

    const data = await loginResponse.json();
    if (data.token !== undefined) {
      cookies().set({
        name: "session",
        value: data.token,
        httpOnly: true,
        secure: true,
        expires: new Date(data.expires),
        path: "/",
      });

      cookies().set({
        name: "role",
        value: data.role,
        httpOnly: true,
        secure: true,
        expires: new Date(data.expires),
        path: "/",
      });
    }
  } catch (e) {
    console.log("Failure to log user in:", e);
  }

  if (
    cookies().get("session") !== undefined &&
    cookies().get("session")?.value !== undefined &&
    cookies().get("session")?.value !== ""
  ) {
    redirect(`/restaurants`);
  }
};

export const logoutUser = async (formData: FormData) => {
  try {
    cookies().delete("session");
    cookies().delete("role");
  } catch (e) {
    console.log("Failure to log user in:", e);
  }

  if (
    cookies().get("session") === undefined ||
    cookies().get("session")?.value !== undefined ||
    cookies().get("session")?.value !== ""
  ) {
    redirect(`/`);
  }
};
