import { useState } from "react";

export default useApi = (apiFunc) => {
  const [data, setData] = useState([]);
  const [error, setError] = useState(false);
  const [isLoading, setIsLoading] = useState(false);

  const request = async (...args) => {
    setIsLoading(true);
    const response = await apiFunc(...args);
    setIsLoading(false);

    if (!response.ok) {
      setError(true);
      return response;
    }

    setError(false);
    setData(response.data);
    return response;
  };

  return { request, data, error, isLoading };
};
