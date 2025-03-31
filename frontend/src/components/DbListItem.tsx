interface Props {
    title: string
    onClick?: any
}
export default function DbListItem({ title, onClick }: Props) {
    return (
        <div
            className="w-full mx-auto p-3 text-center cursor-pointer bg-gray-500 rounded"
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
