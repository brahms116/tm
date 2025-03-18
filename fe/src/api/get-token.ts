export const getToken = (): string | undefined => {
  return localStorage.getItem('api-key') || undefined;  
}
