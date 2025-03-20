import { TimelineResponseItem } from "@/contracts";
import { TransactionTimeline, TransactionTimelineDataItem } from "./timeline";
import { TransactionsTable } from "./transactions-table";
import { MonthSelect } from "./month-select";
import { CategorySelect } from "./category-select";
import { Button } from "@/components/ui/button";
import { Upload } from "lucide-react";
import { Dialog, DialogTrigger } from "@/components/ui/dialog";
import { FileUploadDialog } from "./file-upload-dialog";
import { useState } from "react";
import { startOfMonth, subMonths } from "date-fns";
import { UTCDate } from "@date-fns/utc";
import { useReportTimelineQuery } from "@/api";

const utcDateToDate = (d: UTCDate): Date => {
  return new Date(d.toISOString());
};

const resToTimelineItems = (
  ts: TimelineResponseItem[]
): TransactionTimelineDataItem[] =>
  [...prev12Months].reverse().map((m) => {
    const i = ts.find((t) => {
      return new Date(t.month).toISOString() === m.toISOString();
    });
    return i
      ? {
          month: m,
          spending: i.summary.spendingCents / 100,
          earning: i.summary.earningCents / 100,
        }
      : {
          month: m,
        };
  });

const thisMonth = utcDateToDate(startOfMonth(new UTCDate()));

const prev12Months = Array.from({ length: 12 }, (_, i) => {
  return subMonths(thisMonth, i + 1);
});

export const DashboardPage: React.FC = () => {
  const [rm, setRm] = useState(prev12Months[0]);

  const timelineQuery = useReportTimelineQuery({
    startDate: prev12Months[11].toISOString(),
    endDate: prev12Months[0].toISOString(),
  });

  const chartItems = resToTimelineItems(timelineQuery.data?.items ?? []);

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
      <TransactionTimeline data={chartItems} />
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
