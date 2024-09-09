import { uploadRestaurants } from "@/server/admin";

export default async function Admin() {
  return (
    <main className="w-full flex items-center justify-center">
      <form
        action={uploadRestaurants}
        encType="multipart/form-data"
        className="flex flex-col justify-center items-center"
      >
        <div className="bg-gray-100 rounded-lg p-2 mb-6 flex flex-col gap-6 items-start justify-between">
          <h1>Restaurant File Upload</h1>
          <input type="file" name="file" />
        </div>
        <button type="submit" className="bg-blue-300 rounded-lg px-6 py-2">
          Upload
        </button>
      </form>
    </main>
  );
}
