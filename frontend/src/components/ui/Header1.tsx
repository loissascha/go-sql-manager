interface Props {
    children: any
}
export default function ({ children }: Props) {
    return <h1 className="text-xl mb-3">{children}</h1>
}
