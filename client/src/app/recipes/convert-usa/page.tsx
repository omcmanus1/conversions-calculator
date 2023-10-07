"use client";

import { singleInput, singleOutput } from "@/types/conversionTypes";
import { useState } from "react";
import SingleConversion from "../components/SingleConversion";

export default function Page(): JSX.Element {
  const [input, setInput] = useState<singleInput>({
    ingredient: "",
    inputSystem: "usa",
    inputUnit: "",
    outputSystem: "metric",
    outputUnit: "",
    type: "",
    amount: 0,
  });
  const [output, setOutput] = useState<singleOutput>({
    ingredient: "",
    unit: "",
    amount: 0,
  });

  const volumeInputs = ["cups", "gallons", "quarts", "pints", "fluid oz"];
  const weightInputs = ["cups", "lbs", "oz"];
  const volumeOutputs = ["millilitres", "litres"];
  const weightOutputs = ["grams", "kg"];

  return (
    <SingleConversion
      input={input}
      setInput={setInput}
      output={output}
      setOutput={setOutput}
      weightInputs={weightInputs}
      volumeInputs={volumeInputs}
      weightOutputs={weightOutputs}
      volumeOutputs={volumeOutputs}
    />
  );
}
