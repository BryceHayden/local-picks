import { Suspense } from "react";
import { getRestaurantDetails } from "@/server/restaurants";

export default async function RestaurantDetails({ params }: any) {
  const details = await getRestaurantDetails(params.id);

  return (
    <main className="w-full flex items-center justify-center text-2xl">
      <Suspense fallback={<span>Loading...</span>}>
        <div>
          <h1>{details.Name}</h1>
          {details.Hours.map((info, index) => (
            <div key={index}>
              {info.Day}: {info.Opening} - {info.Closing}
            </div>
          ))}
        </div>
      </Suspense>
    </main>
  );
}
