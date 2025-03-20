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
import { addMonths, isEqual, startOfMonth, subMonths } from "date-fns";
import { useReportQuery, useReportTimelineQuery } from "@/api";
import { utcEquiv } from "@/date-utils";
import { Cents } from "@/components/cents";

const resToTimelineItems = (
  ts: TimelineResponseItem[]
): TransactionTimelineDataItem[] =>
  [...prev12Months].reverse().map((m) => {
    const i = ts.find((t) => {
      return isEqual(new Date(t.month), utcEquiv(m));
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

const thisMonth = startOfMonth(new Date());

const prev12Months = Array.from({ length: 12 }, (_, i) => {
  return subMonths(thisMonth, i + 1);
});

export const DashboardPage: React.FC = () => {
  const [rm, setRm] = useState(prev12Months[0]);
  const [u100, setU100] = useState(false);
  const [isUploadDialogOpen, setIsUploadDialogOpen] = useState(false);

  const timelineQuery = useReportTimelineQuery({
    startDate: utcEquiv(prev12Months[11]).toISOString(),
    endDate: utcEquiv(thisMonth).toISOString(),
  });

  const chartItems = resToTimelineItems(timelineQuery.data?.items ?? []);

  const reportQuery = useReportQuery({
    endDate: utcEquiv(addMonths(rm, 1)).toISOString(),
    startDate: utcEquiv(rm).toISOString(),
    u100,
  });

  return (
    <div className="w-full p-16">
      <div className="mb-12 sm:flex justify-between items-center">
        <h1 className="sm:mb-0 mb-4 text-4xl font-extrabold">Dashboard</h1>
        <Dialog open={isUploadDialogOpen} onOpenChange={setIsUploadDialogOpen}>
          <DialogTrigger asChild>
            <Button className="">
              <Upload size={16} />
              Upload transactions
            </Button>
          </DialogTrigger>
          <FileUploadDialog onClose={() => setIsUploadDialogOpen(false)} />
        </Dialog>
      </div>
      <TransactionTimeline data={chartItems} />
      <div className="mb-3 mt-12">
        <MonthSelect options={prev12Months} value={rm} onChange={setRm} />
      </div>
      <div className="mb-6">
        <CategorySelect isU100={u100} setIsU100={setU100} />
      </div>
      <div>
        <SummaryItem
          className="py-4"
          heading="Total Spending"
          cents={reportQuery.data?.summary.spendingCents ?? 0}
          type="negative"
        />
        <SummaryItem
          className="py-4"
          heading="Total Earning"
          cents={reportQuery.data?.summary.earningCents ?? 0}
        />
        <SummaryItem
          className="py-4"
          heading="Net"
          cents={reportQuery.data?.summary.netCents ?? 0}
        />
      </div>
      <h3 className="mb-6 mt-12 text-2xl font-semibold">Top Spendings</h3>
      <TransactionsTable data={reportQuery.data?.topSpendings ?? []} />
      <h3 className="mb-6 mt-12 text-2xl font-semibold">Top Earnings</h3>
      <TransactionsTable data={reportQuery.data?.topEarnings ?? []} />
    </div>
  );
};

export const SummaryItem: React.FC<{
  heading: string;
  cents: number;
  type?: "negative" | "positive";
  className?: string;
}> = ({ heading, cents, className, type }) => {
  return (
    <div className={className}>
      <h4 className="scroll-m-20 mb-1 text-ml font-semibold tracking-tight">
        {heading}:
      </h4>
      <Cents cents={cents} type={type} />
    </div>
  );
};
