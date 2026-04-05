import { ComponentPropsWithRef } from "react";

type FormLineProps = {
  type : 'name' | 'email' | 'password' | 'expense' | 'amount';
  inputProps?: ComponentPropsWithRef<"input">;
}

const FIELD_CONFIG = {
  name: {
    label: 'name',
    inputType: 'text',
    placeholder: '名前を入力してください',
    autoComplete: 'name',
  },
  email: {
    label: 'メールアドレス',
    inputType: 'email',
    placeholder: 'メールアドレスを入力してください',
    autoComplete: 'email',
  },
  password: {
    label: 'パスワード',
    inputType: 'password',
    placeholder: 'パスワードを入力してください',
    autoComplete: 'current-password',
  },
  expense: {
    label: '支出名',
    inputType: 'text',
    placeholder: '支出を入力してください',
    autoComplete: 'off',
  },
  amount: {
    label: '金額',
    inputType: 'number',
    placeholder: '金額を入力してください',
    autoComplete: 'off',
  },
} as const;

export const FormLine = ({ type, inputProps }: FormLineProps) => {
  const config = FIELD_CONFIG[type]

  return (
    <div className="flex flex-col gap-1">
      <label htmlFor={type} className="text-sm font-medium">
        {config.label}
      </label>
      <input
        id={type}
        name={type}
        type={config.inputType}
        placeholder={config.placeholder}
        autoComplete={config.autoComplete}
        required
        className="border p-2 rounded"
        {...inputProps}
      />
    </div>
  );
}