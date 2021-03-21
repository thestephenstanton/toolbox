import React from "react";
import { useFormikContext } from "formik";

import ImageInputList from "../ImageInputList";
import ErrorMessage from "./ErrorMessage";

export default function FormImagePicker({ name }) {
  const { setFieldValue, errors, touched, values } = useFormikContext();

  const imageUris = values[name];

  const handleAdd = (newUri) => {
    setFieldValue(name, [...imageUris, newUri]);
  };

  const handleRemove = (uri) => {
    setFieldValue(
      name,
      imageUris.filter((u) => u != uri)
    );
  };

  return (
    <>
      <ImageInputList
        imageUris={imageUris}
        onAddImage={handleAdd}
        onRemoveImage={handleRemove}
      />

      <ErrorMessage error={errors[name]} isVisible={touched[name]} />
    </>
  );
}
