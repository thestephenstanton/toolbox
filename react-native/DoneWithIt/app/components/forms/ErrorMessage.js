import { HTML5_FMT } from "moment";
import React from "react";
import { StyleSheet } from "react-native";
import Text from "../Text";

export default function ErrorMessage({ error, isVisible }) {
  if (!isVisible || !error) return null;

  return <Text style={styles.error}>{error}</Text>;
}

const styles = StyleSheet.create({
  error: {
    color: "red",
  },
});
