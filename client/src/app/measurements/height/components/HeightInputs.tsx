import SelectSh from "@/components/select";
import { Input } from "@/components/ui/input";
import { HeightFeet, HeightMetric, ValidInputs } from "@/types/heightTypes";
import { Dispatch, SetStateAction } from "react";

type Props = {
  input: ValidInputs | "";
  setInput: Dispatch<SetStateAction<ValidInputs | "">>;
  setHeightFeet: Dispatch<SetStateAction<HeightFeet>>;
  setHeightMetric: Dispatch<SetStateAction<HeightMetric>>;
};

export default function HeightInputs({
  input,
  setInput,
  setHeightFeet,
  setHeightMetric,
}: Props) {
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
            onChange={(e) =>
              setHeightFeet((prev) => ({ ...prev, feet: Number(e.target.value) }))
            }
          />
          <Input
            type="number"
            max={11}
            placeholder="Inches"
            onChange={(e) =>
              setHeightFeet((prev) => ({ ...prev, inches: Number(e.target.value) }))
            }
          />
        </>
      )}
    </div>
  );
}
