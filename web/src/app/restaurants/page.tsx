"use client";

import { Suspense, useActionState } from "react";
import { getRestaurants } from "@/server/restaurants";
import Link from "next/link";

export default function Restaurants() {
  const [restaurants, formAction] = useActionState(getRestaurants, []);

  const dayOptions = [
    { value: "", label: "Any" },
    { value: "Mon", label: "Monday" },
    { value: "Tue", label: "Tuesday" },
    { value: "Wed", label: "Wednesday" },
    { value: "Thu", label: "Thursday" },
    { value: "Fri", label: "Friday" },
    { value: "Sat", label: "Saturday" },
    { value: "Sun", label: "Sunday" },
  ];

  const timeOptions = [{ value: "", label: "Any" }];
  for (let i = 0; i < 24; i++) {
    timeOptions.push({ value: `${i}:00`, label: `${i}:00` });
  }

  return (
    <main className="w-full flex flex-col gap-10 items-center justify-center">
      <form
        action={formAction}
        className="flex flex-col items-center justify-center gap-8"
      >
        <div className="flex items-center justify-center gap-8 bg-gray-100 rounded-lg p-10">
          <div>
            <label htmlFor="day">Select a day:</label>
            <select name="day" id="day">
              {dayOptions.map((day, index) => (
                <option key={index} value={day.value}>
                  {day.label}
                </option>
              ))}
            </select>
          </div>
          <div>
            <label htmlFor="time">Select a time:</label>
            <select name="time" id="time">
              {timeOptions.map((time, index) => (
                <option key={index} value={time.value}>
                  {time.label}
                </option>
              ))}
            </select>
          </div>
        </div>
        <button type="submit" className="bg-blue-300 rounded-lg px-6 py-2">
          Search
        </button>
      </form>
      <Suspense fallback={<span>Loading...</span>}>
        {restaurants.length === 0 ? (
          <></>
        ) : (
          <>
            <div className="flex flex-col w-full gap-6">
              <div className="self-start text-3xl">Restaurants:</div>
              <div className="self-end flex flex-col w-full items-center text-2xl">
                {restaurants.map((r, index) => (
                  <Link key={index} href={`/restaurants/${r.ID}`}>
                    {r.Name}
                  </Link>
                ))}
              </div>
            </div>
          </>
        )}
      </Suspense>
    </main>
  );
}
