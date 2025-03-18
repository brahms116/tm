interface Configuration {
  apiUrl: string;
}

const mode = import.meta.env.MODE;

const cs: Record<string, Configuration> = {
  development: {
    apiUrl: "http://localhost:8081",
  },
  production: {
    apiUrl: "",
  },
};

const getConfiguration = (): Configuration => {
  return cs[mode];
}

export default getConfiguration();
