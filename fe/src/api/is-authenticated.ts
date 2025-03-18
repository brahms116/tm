import { post } from "./post";

export const isAuthenticated = async (): Promise<boolean> => {
  try {
    await post("/is-authenticated", {});
    return true;
  } catch (error) {
    return false;
  }
};
