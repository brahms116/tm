import {
  Select,
  SelectContent,
  SelectItem,
  SelectTrigger,
  SelectValue,
} from "@/components/ui/select";
import { format } from "date-fns";

export const MonthSelect: React.FC<{
  value: Date;
  options: Date[];
  onChange: (value: Date) => void;
}> = ({ value, options, onChange }) => {

  const fmtD = (d: Date) => format(d, "MMM yyyy");

  console.log('options', options)

  return (
    <Select
      value={value.toISOString()}
      onValueChange={(d) => onChange(new Date(d))}
    >
      <SelectTrigger className="w-[180px]">
        <SelectValue placeholder="Theme">{fmtD(value)}</SelectValue>
      </SelectTrigger>
      <SelectContent>
        {options.map((d) => (
          <SelectItem value={d.toISOString()} key={d.toISOString()}>
            {fmtD(d)}
          </SelectItem>
        ))}
      </SelectContent>
    </Select>
  );
};
