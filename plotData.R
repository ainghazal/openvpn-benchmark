library(jsonlite)
library(ggplot2)

df1 <- fromJSON(txt='openvpn.json')
p <- ggplot(df1, aes(x=loss, y=elapsed)) + geom_point()
ggsave("scatterplot-openvpn.png", plot = p, width = 6, height = 4, dpi = 300)
