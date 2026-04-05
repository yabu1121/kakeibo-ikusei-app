type FormLineProps = {
  type : 'name' | 'email' | 'password';
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
  }
}

export const FormLine = ({ type }: FormLineProps) => {
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
      />
    </div>
  );
}