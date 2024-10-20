import { Dispatch, SetStateAction } from "react";
import { Input } from "@/components/ui/input";
import { BodyWeightTypes, validPaths } from "@/types/bodyWeightTypes";
import TotalInputs from "./TotalInputs";

type Props = {
  bodyWeights: BodyWeightTypes;
  setBodyWeights: Dispatch<SetStateAction<BodyWeightTypes>>;
  setEndpointPath: Dispatch<SetStateAction<validPaths>>;
  converted: boolean;
};

export default function BodyWeightInputs({
  bodyWeights,
  setBodyWeights,
  setEndpointPath,
  converted,
}: Props) {
  return (
    <div className="flex flex-col gap-1 start-0">
      <TotalInputs
        bodyWeights={bodyWeights}
        setBodyWeights={setBodyWeights}
        setEndpointPath={setEndpointPath}
        converted={converted}
      />
      <div className="flex flex-row gap-1">
        <div className="flex flex-col">
          <h2 className="text-left p-1">Stone</h2>
          <Input
            type="number"
            placeholder="0"
            readOnly={converted}
            onChange={(e) => {
              setBodyWeights({ ...bodyWeights, stone: Number(e.target.value) });
              setEndpointPath("bodyweight-stone");
            }}
            value={!bodyWeights.stone ? "" : bodyWeights.stone}
          />
        </div>
        <div className="flex flex-col">
          <h2 className="text-left p-1">Lbs</h2>
          <Input
            type="number"
            placeholder="0"
            readOnly={converted}
            onChange={(e) => {
              setBodyWeights({ ...bodyWeights, lbs: Number(e.target.value) });
              setEndpointPath("bodyweight-stone");
            }}
            value={!bodyWeights.lbs ? "" : bodyWeights.lbs}
          />
        </div>
      </div>
    </div>
  );
}
