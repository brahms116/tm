import { TimelineResponseItem } from "@/contracts";
import { Graph } from "./graph";
import { TransactionsTable } from "./transactions-table";
import { MonthSelect } from "./month-select";
import { CategorySelect } from "./category-select";
import { Button } from "@/components/ui/button";
import { Upload } from "lucide-react";
import { Dialog, DialogTrigger } from "@/components/ui/dialog";
import { FileUploadDialog } from "./file-upload-dialog";
import { useState } from "react";
import { startOfMonth, subMonths } from "date-fns";
import { useReportTimelineQuery } from "@/api";

const demoData: TimelineResponseItem[] = [
  {
    month: "2023-12-01T00:00:00.000Z",
    summary: {
      spendingCents: 3274,
      earningCents: 7312,
      netCents: 4038,
    },
  },
  {
    month: "2024-01-01T00:00:00.000Z",
    summary: {
      spendingCents: 2187,
      earningCents: 6745,
      netCents: 4558,
    },
  },
  {
    month: "2024-02-01T00:00:00.000Z",
    summary: {
      spendingCents: 4193,
      earningCents: 8102,
      netCents: 3909,
    },
  },
  {
    month: "2024-03-01T00:00:00.000Z",
    summary: {
      spendingCents: 3786,
      earningCents: 8914,
      netCents: 5128,
    },
  },
  {
    month: "2024-04-01T00:00:00.000Z",
    summary: {
      spendingCents: 2511,
      earningCents: 6923,
      netCents: 4412,
    },
  },
  {
    month: "2024-05-01T00:00:00.000Z",
    summary: {
      spendingCents: 4532,
      earningCents: 9123,
      netCents: 4591,
    },
  },
  {
    month: "2024-06-01T00:00:00.000Z",
    summary: {
      spendingCents: 3924,
      earningCents: 7987,
      netCents: 4063,
    },
  },
  {
    month: "2024-07-01T00:00:00.000Z",
    summary: {
      spendingCents: 2675,
      earningCents: 7054,
      netCents: 4379,
    },
  },
  {
    month: "2024-08-01T00:00:00.000Z",
    summary: {
      spendingCents: 3821,
      earningCents: 8293,
      netCents: 4472,
    },
  },
  {
    month: "2024-09-01T00:00:00.000Z",
    summary: {
      spendingCents: 3120,
      earningCents: 7882,
      netCents: 4762,
    },
  },
  {
    month: "2024-10-01T00:00:00.000Z",
    summary: {
      spendingCents: 4312,
      earningCents: 9011,
      netCents: 4699,
    },
  },
  {
    month: "2024-11-01T00:00:00.000Z",
    summary: {
      spendingCents: 2987,
      earningCents: 7423,
      netCents: 4436,
    },
  },
];

const timelineResponseItemsToChartData = (is: TimelineResponseItem[]) => {}

const thisMonth = startOfMonth(new Date());

const prev12Months = Array.from({ length: 12 }, (_, i) => {
  return subMonths(thisMonth, i + 1);
});

export const DashboardPage: React.FC = () => {
  const [rm, setRm] = useState(prev12Months[0]);

  const timelineQuery = useReportTimelineQuery({
    startDate: prev12Months[11].toISOString(),
    endDate: prev12Months[0].toISOString(),
  });

  return (
    <div className="w-full p-16">
      <div className="mb-12 sm:flex justify-between items-center">
        <h1 className="sm:mb-0 mb-4 text-4xl font-extrabold">Dashboard</h1>
        <Dialog>
          <DialogTrigger asChild>
            <Button className="">
              <Upload size={16} />
              Upload transactions
            </Button>
          </DialogTrigger>
          <FileUploadDialog />
        </Dialog>
      </div>
      <Graph data={demoData} />
      <div className="mb-3 mt-12">
        <MonthSelect options={prev12Months} value={rm} onChange={setRm} />
      </div>
      <div className="mb-6">
        <CategorySelect />
      </div>
      <div>
        <div className="py-4">
          <h4 className="scroll-m-20 mb-1 text-ml font-semibold tracking-tight">
            Total Spending:
          </h4>
          <p className="text-red-400">$2987</p>
        </div>
        <div className="py-4">
          <h4 className="scroll-m-20 mb-1 text-ml font-semibold tracking-tight">
            Total Earning:
          </h4>
          <p className="text-lime-500">$2300</p>
        </div>
        <div className="py-4">
          <h4 className="scroll-m-20 mb-1 text-ml font-semibold tracking-tight">
            Net:
          </h4>
          <p className="text-lime-500">$2300</p>
        </div>
      </div>
      <h3 className="mb-6 mt-12 text-2xl font-semibold">Top Spendings</h3>
      <TransactionsTable />
      <h3 className="mb-6 mt-12 text-2xl font-semibold">Top Earnings</h3>
      <TransactionsTable />
    </div>
  );
};

