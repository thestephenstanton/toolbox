import * as SecureStore from "expo-secure-store";
import jwtDecode from "jwt-decode";

const key = "jwt";
const storeJwt = async (jwt) => {
  try {
    await SecureStore.setItemAsync(key, jwt);
  } catch (error) {
    console.log("error storing jwt", error);
  }
};

const getJwt = async () => {
  try {
    return await SecureStore.getItemAsync(key);
  } catch (error) {
    console.log("error getting jwt", error);
  }
};

const getUser = async () => {
  const jwt = await getJwt();
  if (!jwt) {
    return null;
  }

  return jwtDecode(jwt);
};

const removeJwt = async () => {
  try {
    await SecureStore.deleteItemAsync(key);
  } catch (error) {
    console.log("error removing jwt", error);
  }
};

export default {
  getUser,
  getJwt,
  storeJwt,
  removeJwt,
};
