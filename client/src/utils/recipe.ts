import { singleInput } from "@/types/conversionTypes";
import { Dispatch, SetStateAction } from "react";

export const handleInput = <K extends keyof singleInput>(
  setInput: Dispatch<SetStateAction<singleInput>>,
  input: singleInput,
  property: K,
  value: singleInput[K]
) => {
  setInput({
    ...input,
    [property]: value === "string" ? value.toLowerCase() : value,
  });
};

export const inputComplete = (input: singleInput): boolean => {
  return Object.values(input).every((item) => !!item);
};
