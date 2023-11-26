import { RecipeInput } from "@/types/conversionTypes";

export const inputComplete = (input: RecipeInput): boolean => {
  if (input.amount < 0) return false;
  return Object.values(input).every((item) => !!item);
};
