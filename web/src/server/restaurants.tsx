"use server";

import { revalidatePath } from "next/cache";
import { NextRequest } from "next/server";

interface iRestaurant {
  ID: string;
  Name: string;
}

export const getRestaurants = async (
  prevState: any,
  formData: FormData
): Promise<iRestaurant[]> => {
  const req = await fetch("http://api:9000/v1/search", {
    body: JSON.stringify({
      day: formData.get("day"),
      time: formData.get("time"),
    }),
    method: "POST",
  });

  let data = req.status !== 200 ? { restaurants: [] } : await req.json();
  console.log("DATA -----\n\n", data);
  revalidatePath('/restaurants"');
  return data.restaurants as iRestaurant[];
};

interface DailyHours {
  Day: string;
  Opening: string;
  Closing: string;
}

interface iRestaurantDetails {
  ID: string;
  Name: string;
  Hours: Array<DailyHours>;
}

export const getRestaurantDetails = async (
  id: string
): Promise<iRestaurantDetails> => {
  let details = { ID: "Unknown", Name: "Unsure", Hours: [] };
  try {
    const req = await fetch(`http://api:9000/v1/restaurants/${id}`, {
      method: "GET",
    });

    let data = req.status === 404 ? {} : await req.json();
    details = data.restaurant;
  } catch (e) {
    console.log("Error fetching restaurant details", e);
  }

  return details as iRestaurantDetails;
};
