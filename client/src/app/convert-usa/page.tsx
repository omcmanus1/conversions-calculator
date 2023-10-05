"use client";

import SelectSh from "@/components/select";
import { Button } from "@/components/ui/button";
import { Input } from "@/components/ui/input";
import { singleInput } from "@/types/inputTypes";
import { useState } from "react";

export default function Page() {
  const [input, setInput] = useState<singleInput>({
    ingredient: "",
    inputSystem: "usa",
    inputUnit: "",
    outputSystem: "metric",
    outputUnit: "",
    type: "",
    amount: 0,
  });
  const volumeInputs = ["cups", "gallons", "quarts", "pints", "fluid oz"];
  const weightInputs = ["cups", "lbs", "oz"];
  const volumeOutputs = ["millilitres", "litres"];
  const weightOutputs = ["grams", "kg"];

  console.log("🚀 ~ file: page.tsx:19 ~ Page ~ input:", input);

  const handleInput = <K extends keyof singleInput>(
    property: K,
    value: singleInput[K]
  ) => {
    setInput({
      ...input,
      [property]: value === "string" ? value.toLowerCase() : value,
    });
  };

  return (
    <>
      <Input
        className="mb-1"
        type="text"
        placeholder="Ingredient"
        onChange={(e) => handleInput("ingredient", e.target.value)}
      />
      {!!input.ingredient && (
        <SelectSh
          handleChange={(e) => handleInput("type", e)}
          placeholder="Type"
          selectContent={["weight", "volume"]}
        />
      )}
      {!!input.type && (
        <SelectSh
          handleChange={(e) => handleInput("inputUnit", e)}
          placeholder="Input Unit"
          selectContent={
            input.type === "weight"
              ? weightInputs
              : input.type === "volume"
              ? volumeInputs
              : ["Please specify a unit"]
          }
        />
      )}
      {!!input.inputUnit && (
        <Input
          className="mb-1"
          type="number"
          placeholder="Amount"
          onChange={(e) => handleInput("amount", Number(e.target.value))}
        />
      )}
      {!!input.amount && (
        <SelectSh
          handleChange={(e) => handleInput("outputUnit", e)}
          placeholder="Output Unit"
          selectContent={
            input.type === "weight"
              ? weightOutputs
              : input.type === "volume"
              ? volumeOutputs
              : ["Please specify a unit"]
          }
        />
      )}
      <Button className="mt-3 mb-3" disabled variant="outline">
        Convert From Freedom Units
      </Button>
      <Input disabled />
    </>
  );
}
