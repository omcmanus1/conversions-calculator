import { RecipeInput } from "@/types/conversionTypes";

const baseUrl = "http://localhost:8080/api/convert";

export const getRequest = async () => {
  try {
    const res = await fetch(`${baseUrl}/get-encode`);
    if (!res.ok) {
      throw new Error("Failed to fetch conversions");
    }
    const jsonData = await res.json();
    return jsonData[0];
  } catch (e) {
    console.log("getRequest: ", e);
  }
};

export const postRequest = async (
  path: string,
  data: RecipeInput | Array<RecipeInput>
) => {
  try {
    const res = await fetch(`${baseUrl}/${path}`, {
      method: "POST",
      mode: "cors",
      cache: "default",
      credentials: "same-origin",
      headers: {
        "Content-Type": "application/json",
      },
      redirect: "follow",
      body: JSON.stringify(data),
    });
    const jsonData = await res.json();
    return jsonData;
  } catch (e) {
    console.log("postRequest: ", e);
  }
};
