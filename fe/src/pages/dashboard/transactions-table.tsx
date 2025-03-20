import { Transaction } from "@/contracts";
import { ColumnDef } from "@tanstack/react-table";
import { format } from "date-fns";
import { DataTable } from "./data-table";
import { Cents } from "@/components/cents";

export const columns: ColumnDef<Transaction>[] = [
  {
    header: "Date",
    cell: ({ row }) => format(new Date(row.original.date), "MMM d, yyyy"),
  },
  {
    header: "Amount",
    cell: ({ row }) => <Cents cents={row.original.amountCents} />,
  },
  {
    header: "Description",
    cell: ({ row }) => row.original.description,
  },
];

export const TransactionsTable: React.FC<{
  data: Transaction[];
}> = ({ data }) => {
  return <DataTable columns={columns} data={data} />;
};
