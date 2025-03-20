export const Cents: React.FC<{
  cents: number;
  type?: "positive" | "negative";
}> = ({ cents, type }) => {
  const dollars = new Intl.NumberFormat("en-US", {
    style: "currency",
    maximumFractionDigits: 2,
    currency: "USD",
  }).format(cents / 100);

  const dt = type ?? (cents < 0 ? "negative" : "positive");

  return (
    <span className={dt === "negative" ? "text-red-400" : "text-lime-500"}>
      {dollars}
    </span>
  );
};
