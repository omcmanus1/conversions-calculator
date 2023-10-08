"use client";

import SingleConversion from "../components/SingleConversion";

export default function Page(): JSX.Element {
  const volumeInputs = ["cups", "gallons", "quarts", "pints", "fluid oz"];
  const weightInputs = ["cups", "lbs", "oz"];
  const volumeOutputs = ["millilitres", "litres"];
  const weightOutputs = ["grams", "kg"];

  return (
    <SingleConversion
      weightInputs={weightInputs}
      volumeInputs={volumeInputs}
      weightOutputs={weightOutputs}
      volumeOutputs={volumeOutputs}
    />
  );
}
