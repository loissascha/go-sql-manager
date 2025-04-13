interface Props {
    htmlFor?: string
    children?: any
}
export default function FormLabel({ htmlFor, children }: Props) {
    return <label htmlFor={htmlFor}>{children}</label>
}
