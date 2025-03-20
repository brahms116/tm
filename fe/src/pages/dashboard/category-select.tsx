import {
  DropdownMenu,
  DropdownMenuContent,
  DropdownMenuCheckboxItem,
  DropdownMenuTrigger,
} from "@/components/ui/dropdown-menu";
import { Button } from "@/components/ui/button";

export const CategorySelect: React.FC<{
  isU100: boolean;
  setIsU100: (u100: boolean) => void;
}> = ({ isU100, setIsU100 }) => {
  return (
    <DropdownMenu>
      <DropdownMenuTrigger asChild>
        <Button variant="outline">Categories ({isU100 ? "1" : "All"})</Button>
      </DropdownMenuTrigger>
      <DropdownMenuContent>
        <DropdownMenuCheckboxItem
          checked={!isU100}
          onClick={() => setIsU100(false)}
        >
          All
        </DropdownMenuCheckboxItem>
        <DropdownMenuCheckboxItem
          checked={isU100}
          onClick={() => setIsU100(true)}
        >
          Under $100
        </DropdownMenuCheckboxItem>
      </DropdownMenuContent>
    </DropdownMenu>
  );
};
