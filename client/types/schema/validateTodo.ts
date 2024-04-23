import * as yup from "yup";

const validateTodoSchema = yup.object().shape({
  todo: yup
    .string()
    .required("Todo is required")
    .max(30, "Todo must be less than 30 characters"),
});

export default validateTodoSchema;
