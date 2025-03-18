export interface Transaction {
  id: string;
  date: string;
  description: string;
  amountCents: number;
  category?: string;
}
export interface TimelineRequest {
  startDate: string;
  endDate: string;
}
export interface ReportRequest {
  startDate: string;
  endDate: string;
  u100: boolean;
}
export interface TimelineResponseItem {
  month: string;
  summary: Summary;
}
export interface TimelineResponse {
  items: TimelineResponseItem[];
}
export interface ReportResponse {
  summary: Summary;
  topSpendings: Transaction[];
  topEarnings: Transaction[];
}
export interface Summary {
  spendingCents: number;
  earningCents: number;
  netCents: number;
}
