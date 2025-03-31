interface Props {
    title: string
    onClick?: any
}
export default function DbListItem({ title, onClick }: Props) {
    return (
        <div
            className="w-full p-3 text-center cursor-pointer"
            onClick={() => {
                if (onClick) {
                    onClick()
                }
            }}
        >
            {title}
        </div>
    )
}
