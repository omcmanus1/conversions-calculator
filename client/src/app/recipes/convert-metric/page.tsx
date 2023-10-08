"use client";

import SingleConversion from "../components/SingleConversion";

export default function Page() {
  const volumeInputs = ["millilitres", "litres"];
  const weightInputs = ["grams", "kg"];
  const volumeOutputs = ["cups", "gallons", "quarts", "pints", "fluid oz"];
  const weightOutputs = ["cups", "lbs", "oz"];

  return (
    <SingleConversion
      weightInputs={weightInputs}
      volumeInputs={volumeInputs}
      weightOutputs={weightOutputs}
      volumeOutputs={volumeOutputs}
    />
  );
}
