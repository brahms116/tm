import { addMinutes } from "date-fns";

export const utcEquiv = (d: Date): Date =>
  addMinutes(d, -d.getTimezoneOffset());

export const localEquiv = (d: Date): Date =>
  addMinutes(d, d.getTimezoneOffset());
