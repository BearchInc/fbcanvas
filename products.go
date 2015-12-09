package main

type Product struct {
	Image string `json:"image"`
	Name  string `json:"name"`
	Price string `json:"price"`
}

var products = []Product{
	Product{
		Name: "Carrinho Hot Wheels",
		Price: "7.90",
		Image: "http://iacom.s8.com.br/produtos/01/00/item/123923/8/123923825_1GG.jpg",
	},
	Product{
		Name: "Docinhos e Bebidas para a Festa",
		Price: "786.00 ",
		Image: "http://mlb-s1-p.mlstatic.com/brigadeiro-ou-beijinho-cento-22597-MLB20232317909_012015-F.jpg",
	},
	Product{
		Name: "Bola de Futebol",
		Price: "39.49",
		Image: "http://static5.netshoes.net/Produtos/bola-topper-astro-campo/28/D30-0158-028/D30-0158-028_detalhe1.jpg",
	},
	Product{
		Name: "Chuteira",
		Price: "64.90",
		Image: "http://static5.netshoes.net/Produtos/chuteira-adidas-f5-tf-society-infantil/72/132-8886-172/132-8886-172_detalhe1.jpg",
	},
	Product{
		Name: "Cachorro-Quente para a Festa",
		Price: "650.00",
		Image: "http://a57.foxnews.com/global.fncstatic.com/static/managed/img/876/493/etete5646456.jpg?ve=1&tl=1",
	},
	Product{
		Name: "Kit para Chá de Cozinha",
		Price: "29.90",
		Image: "http://images10.tcdn.com.br/img/img_prod/268642/minhas_panelinhas_panela_e_cia_8331_2_20130128094132.jpg",
	},
	Product{
		Name: "Barbie",
		Price: "79.90",
		Image: "http://www.hamleys.com/images/_lib/barbie-rock-star-doll-84671-0-1431688775000.jpg",
	},
	Product{
		Name: "Bicicleta",
		Price: "242.90",
		Image: "http://www.belasdicas.com/img/fotos/bicicleta%20infantil%203%20anos%206.jpg",
	},
	Product{
		Name: "Boneca Peppa Pig",
		Price: "99.00",
		Image: "http://www.angelamagazine.com.br/media/catalog/product/cache/1/image/9df78eab33525d08d6e5fb8d27136e95/b/o/boneca_peppa_pig_multibrink_3.jpg",
	},
	Product{
		Name: "Jogos – Alquimia",
		Price: "89.90",
		Image: "http://mlb-s1-p.mlstatic.com/jogo-alquimia-grow-13-elementos-e-75-experincias-antigo-19664-MLB20175425305_102014-F.jpg",
	},
	Product{
		Name: "Kit Jogos Educativos ",
		Price: "149.90",
		Image: "http://www.psicopedagogavaleria.com.br/site/images/stories/kit%20jogos%20educativos%201.jpg",
	},
	Product{
		Name: "Boneca",
		Price: "54.90",
		Image: "https://tricae-a.akamaihd.net/p/Milk-Boneca-Girl-Jeans-Milk-7283-47306-1.jpg",
	},
	Product{
		Name: "Bonecos Patati e Patata",
		Price: "163.90",
		Image: "http://mlb-s2-p.mlstatic.com/boneco-patati-patata-213801-MLB20404436684_092015-F.jpg",
	},
	Product{
		Name: "Bonecos – Os Vingadores",
		Price: "94.90",
		Image: "https://www.toys.com.br/media/product/cf8/kit-4-bonecos-marvel-avengers-vingadores-30-cm-hasbro-a8f.jpg",
	},
	Product{
		Name: "Lego Clássico",
		Price: "198.90",
		Image: "http://statics.livrariacultura.net.br/products/capas_lg/919/42889919.jpg",
	},
	Product{
		Name: "Aluguel Fla-Flu para Festa",
		Price: "132.00",
		Image: "http://images.centauro.com.br/900x900/77296200/mesa-de-pebolim-klopf-galera-img.jpg",
	},
	Product{
		Name: "Aluguel Cama Elástica para a Festa",
		Price: "227.00",
		Image: "http://www.brinquedosrabisco.com.br/uploads/Cama-Elastica.jpg",
	},
	Product{
		Name: "Aluguel Tobogã para a Festa",
		Price: "350.00",
		Image: "http://www.iclaz.com.br/foto/GG/250297/1-toboga_7837/locacao-de-brinquedos-inflaveis-toboga-inflavel-festa-de-crianca-aluguel-de-brinquedos-inflaveis.jpg",
	},
}