export function formatTimeDifference(publishedAt: string) {
    const publishedDate = new Date(publishedAt);
    const currentDate = new Date();
    const diffInMilliseconds = currentDate.getTime() - publishedDate.getTime();
    const diffInMinutes = diffInMilliseconds / (1000 * 60);
    const diffInHours = diffInMilliseconds / (1000 * 60 * 60);
    const diffInDays = diffInMilliseconds / (1000 * 60 * 60 * 24);

    if (diffInMinutes < 60) {
        return `${Math.round(diffInMinutes)} minutes ago`;
    } else if (diffInHours < 24) {
        return `${Math.round(diffInHours)} hours ago`;
    } else {
        return `${Math.round(diffInDays)} days ago`;
    }
}