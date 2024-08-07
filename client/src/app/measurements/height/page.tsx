"use client";

import SelectSh from "@/components/select";
import { Input } from "@/components/ui/input";
import { HeightFeet, HeightMetric, ValidInputs } from "@/types/heightTypes";
import { useState } from "react";

export default function ConvertHeight() {
  const [input, setInput] = useState<ValidInputs | "">("");
  const [heightMetric, setHeightMetric] = useState<HeightMetric>({
    centimetres: 0,
    metres: 0,
  });
  const [heightFeet, setHeightFeet] = useState<HeightFeet>({ feet: 0, inches: 0 });

  return (
    <div className="text-center w-[200px] flex flex-col gap-1">
      <SelectSh
        handleChange={(e) => setInput(e as ValidInputs)}
        placeholder="Input Unit"
        selectContent={["centimetres", "metres", "feet"]}
      />
      {input === "centimetres" && (
        <Input
          type="number"
          placeholder="Centimetres"
          onChange={(e) =>
            setHeightMetric({ metres: 0, centimetres: Number(e.target.value) })
          }
        />
      )}
      {input === "metres" && (
        <Input
          type="number"
          placeholder="Metres"
          onChange={(e) =>
            setHeightMetric({ centimetres: 0, metres: Number(e.target.value) })
          }
        />
      )}
      {input === "feet" && (
        <>
          <Input
            type="number"
            placeholder="Feet"
            onChange={(e) => setHeightFeet({ inches: 0, feet: Number(e.target.value) })}
          />
          <Input
            type="number"
            max={11}
            placeholder="Inches"
            onChange={(e) => setHeightFeet({ feet: 0, inches: Number(e.target.value) })}
          />
        </>
      )}
    </div>
  );
}
