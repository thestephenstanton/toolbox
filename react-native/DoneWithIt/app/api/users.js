import client from "./client";

const register = (userInfo) => {
  return client.post("/users", userInfo);
};

export default {
  register,
};
