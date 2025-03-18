import { Transaction } from "@/contracts";
import { ColumnDef } from "@tanstack/react-table";
import { format } from "date-fns";
import { DataTable } from "./data-table";

const demoData: Transaction[] = [
  {
    id: "1f2a3b4c",
    date: "2023-12-15T00:00:00.000Z",
    description: "Grocery Store",
    amountCents: -4532,
    category: "Groceries",
  },
  {
    id: "2b3c4d5e",
    date: "2024-01-10T00:00:00.000Z",
    description: "Salary",
    amountCents: 75000,
    category: "Income",
  },
  {
    id: "3c4d5e6f",
    date: "2024-02-05T00:00:00.000Z",
    description: "Restaurant",
    amountCents: -2345,
    category: "Dining",
  },
  {
    id: "4d5e6f7g",
    date: "2024-03-20T00:00:00.000Z",
    description: "Electric Bill",
    amountCents: -6200,
    category: "Utilities",
  },
  {
    id: "5e6f7g8h",
    date: "2024-04-25T00:00:00.000Z",
    description: "Freelance Payment",
    amountCents: 15000,
    category: "Income",
  },
  {
    id: "6f7g8h9i",
    date: "2024-05-12T00:00:00.000Z",
    description: "Online Subscription",
    amountCents: -999,
    category: "Entertainment",
  },
  {
    id: "7g8h9i0j",
    date: "2024-06-08T00:00:00.000Z",
    description: "Fuel Station",
    amountCents: -4321,
    category: "Transportation",
  },
  {
    id: "8h9i0j1k",
    date: "2024-07-14T00:00:00.000Z",
    description: "Shopping Mall",
    amountCents: -7850,
    category: "Shopping",
  },
  {
    id: "9i0j1k2l",
    date: "2024-08-17T00:00:00.000Z",
    description: "Health Insurance",
    amountCents: -12500,
    category: "Health",
  },
  {
    id: "0j1k2l3m",
    date: "2024-09-22T00:00:00.000Z",
    description: "Bonus Payment",
    amountCents: 20000,
    category: "Income",
  },
  {
    id: "1k2l3m4n",
    date: "2024-10-18T00:00:00.000Z",
    description: "Home Rent",
    amountCents: -35000,
    category: "Housing",
  },
  {
    id: "2l3m4n5o",
    date: "2024-11-05T00:00:00.000Z",
    description: "Gym Membership",
    amountCents: -4500,
    category: "Health",
  },
];

export const columns: ColumnDef<Transaction>[] = [
  {
    header: "Date",
    cell: ({ row }) => format(new Date(row.original.date), "MMM d, yyyy"),
  },
  {
    header: "Amount",
    cell: ({ row }) => `$${(row.original.amountCents / 100).toFixed(2)}`,
  },
  {
    header: "Description",
    cell: ({ row }) => row.original.description,
  },
];

export const TransactionsTable: React.FC = () => {
  return <DataTable columns={columns} data={demoData} />;
};
