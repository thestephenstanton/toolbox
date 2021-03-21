import React from "react";
import { useFormikContext } from "formik";

import Picker from "../Picker";
import ErrorMessage from "./ErrorMessage";

export default function FormPicker({
  name,
  width,
  numColumns,
  PickerItemComponent,
  ...otherProps
}) {
  const { setFieldValue, errors, touched, values } = useFormikContext();

  return (
    <>
      <Picker
        onSelectItem={(item) => setFieldValue(name, item)}
        selectedItem={values[name]}
        width={width}
        numColumns={numColumns}
        PickerItemComponent={PickerItemComponent}
        {...otherProps}
      />
      <ErrorMessage error={errors[name]} isVisible={touched[name]} />
    </>
  );
}
