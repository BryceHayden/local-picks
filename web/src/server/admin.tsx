"use server";

export const uploadRestaurants = async (formData: FormData) => {
  try {
    // body: JSON.stringify({
    //   file: formData.get("file"),
    // }),

    const loginResponse = await fetch("http://api:9000/v1/admin/restaurants", {
      body: formData,
      method: "POST",
    });

    console.log("CHECKING FORM DATA", formData.get("file"), loginResponse);
  } catch (e) {
    console.log("Failure to log user in:", e);
  }
};
