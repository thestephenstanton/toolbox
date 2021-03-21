import { useNavigation } from "@react-navigation/native";
import React, { useEffect, useState } from "react";
import { StyleSheet, FlatList } from "react-native";

import listingsApi from "../api/listings";
import useApi from "../hooks/useApi";
import ActivityIndicator from "../components/ActivityIndicator";
import Button from "../components/Button";
import Card from "../components/Card";
import Screen from "../components/Screen";
import Text from "../components/Text";
import colors from "../config/colors";
import routes from "../navigation/routes";

export default function ListingsScreen() {
  const navigation = useNavigation();

  const { data: listings, error, isLoading, request: loadListings } = useApi(
    listingsApi.getListings
  );

  useEffect(() => {
    loadListings();
  }, []);

  return (
    <>
      <ActivityIndicator isVisible={isLoading} />
      <Screen style={styles.container}>
        {error && (
          <>
            <Text>Couldn't retrieve the listings.</Text>
            <Button title="Retry" onPress={loadListings} />
          </>
        )}
        <FlatList
          style={styles.cardContainer}
          data={listings}
          keyExtractor={(listing) => listing.id.toString()}
          renderItem={({ item }) => (
            <Card
              title={item.title}
              subTitle={"$" + item.price}
              thumbnailUrl={item.images[0].thumbnailUrl}
              imageUrl={item.images[0].url}
              onPress={() => navigation.navigate(routes.LISTING_DETAILS, item)}
            />
          )}
        />
      </Screen>
    </>
  );
}

const styles = StyleSheet.create({
  container: {
    flex: 1,
    backgroundColor: colors.light,
  },
  cardContainer: {
    margin: 15,
  },
});
