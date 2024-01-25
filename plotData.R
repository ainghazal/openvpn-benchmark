library(jsonlite)
library(ggplot2)

df1 <- fromJSON(txt='comparison.json')
p <- ggplot(df1, aes(x=loss, y=elapsed, color=flavor)) + geom_jitter(alpha=0.8, size=3) 
# + geom_point(size=3, alpha=0.8)
ggsave("scatterplot-openvpn.png", plot = p, width = 20, height = 6, dpi = 600)
