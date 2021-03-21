import { create } from "apisauce";
import cache from "../utility/cache";
import authStorage from "../auth/storage";

const apiClient = create({
  baseURL: "http://192.168.1.137:9000/api",
});

apiClient.addAsyncRequestTransform(async (request) => {
  const jwt = await authStorage.getJwt();
  if (!jwt) {
    return;
  }

  request.headers["x-auth-token"] = jwt;
});

const get = apiClient.get;
apiClient.get = async (url, params, axiosConfig) => {
  const response = await get(url, params, axiosConfig);

  if (!response.ok) {
    const data = await cache.get(url);

    return data ? { ok: true, data } : response;
  }

  cache.store(url, response.data);

  return response;
};

export default apiClient;
