import { RecipeInput } from "@/types/conversionTypes";

const devUrl = process.env.NEXT_PUBLIC_DEV_API_URL;
const productionUrl = process.env.NEXT_PUBLIC_PROD_API_URL;

const isProduction = process.env.NODE_ENV === "production";
const apiUrl = isProduction ? productionUrl : devUrl;

export const getRequest = async () => {
  try {
    const res = await fetch(`${apiUrl}/get-encode`);
    if (!res.ok) {
      throw new Error("Failed to fetch conversions");
    }
    const jsonData = await res.json();
    return jsonData[0];
  } catch (e) {
    console.log("getRequest: ", e);
  }
};

export const postRequest = async (path: string, data: RecipeInput | RecipeInput[]) => {
  try {
    const res = await fetch(`${apiUrl}/${path}`, {
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
