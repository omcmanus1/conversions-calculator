"use client";

import SelectSh from "@/components/select";
import { ValidInputs } from "@/types/heightTypes";
import { useState } from "react";

export default function ConvertHeight() {
  const [input, setInput] = useState<keyof typeof ValidInputs | "">("");

  const handleInput = (e: keyof typeof ValidInputs) => {
    setInput(e);
  };

  return (
    <div className="text-center w-[200px]">
      <SelectSh
        handleChange={(e) => handleInput(e as keyof typeof ValidInputs)}
        placeholder="Input Unit"
        selectContent={Object.values(ValidInputs)}
      />
    </div>
  );
}
