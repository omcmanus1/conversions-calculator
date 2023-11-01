import { SingleInput } from "@/types/conversionTypes";
import { Dispatch, SetStateAction } from "react";

export const handleInput = <K extends keyof SingleInput>(
  setInput: Dispatch<SetStateAction<SingleInput>>,
  input: SingleInput,
  property: K,
  value: SingleInput[K]
) => {
  setInput({
    ...input,
    [property]: value === "string" ? value.toLowerCase() : value,
  });
};

export const inputComplete = (input: SingleInput): boolean => {
  return Object.values(input).every((item) => !!item);
};
