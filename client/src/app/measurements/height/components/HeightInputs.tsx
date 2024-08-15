import SelectSh from "@/components/select";
import { Input } from "@/components/ui/input";
import { HeightFeet, HeightMetric, ValidInputs } from "@/types/heightTypes";
import { inputComplete } from "@/utils/heights";
import { Dispatch, SetStateAction } from "react";

type Props = {
  input: ValidInputs | "";
  setInput: Dispatch<SetStateAction<ValidInputs | "">>;
  heightFeet: HeightFeet;
  setHeightFeet: Dispatch<SetStateAction<HeightFeet>>;
  heightMetric: HeightMetric;
  setHeightMetric: Dispatch<SetStateAction<HeightMetric>>;
};

export default function HeightInputs({
  input,
  setInput,
  heightFeet,
  setHeightFeet,
  heightMetric,
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
          className={
            !inputComplete(input, heightFeet, heightMetric)
              ? "border-2 border-red-500"
              : ""
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
          className={
            !inputComplete(input, heightFeet, heightMetric)
              ? "border-2 border-red-500"
              : ""
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
            className={
              !inputComplete(input, heightFeet, heightMetric)
                ? "border-2 border-red-500"
                : ""
            }
          />
          <Input
            type="number"
            min={0}
            max={11}
            placeholder="Inches"
            onChange={(e) =>
              setHeightFeet((prev) => ({ ...prev, inches: Number(e.target.value) }))
            }
            className={
              !inputComplete(input, heightFeet, heightMetric)
                ? "border-2 border-red-500"
                : ""
            }
          />
        </>
      )}
    </div>
  );
}
