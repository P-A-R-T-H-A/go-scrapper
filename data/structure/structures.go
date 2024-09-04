package structure

type BreadCrumb struct {
	Label       string `json:"label"`
	QueryParams []struct {
		Key   string `json:"key"`
		Value string `json:"value"`
	} `json:"queryParams"`
	SearchURL string `json:"searchUrl"`
}

type ProductDetails struct {
	Article struct {
		AddToCartButton struct {
			Description any    `json:"description"`
			Enabled     bool   `json:"enabled"`
			Label       string `json:"label"`
			Link        any    `json:"link"`
			Popup       any    `json:"popup"`
			Title       any    `json:"title"`
		} `json:"addToCartButton"`
		ArticleCode string `json:"articleCode"`
		Attributes  struct {
			Athlete any    `json:"athlete"`
			Color   string `json:"color"`
			Country string `json:"country"`
		} `json:"attributes"`
		Badge struct {
			Color string `json:"color"`
			Text  string `json:"text"`
		} `json:"badge"`
		Badges []struct {
			Color string `json:"color"`
			Text  string `json:"text"`
			Type  string `json:"type"`
		} `json:"badges"`
		Coordinates any `json:"coordinates"`
		Customize   any `json:"customize"`
		DataLayer   struct {
			MasterItemStatus  int      `json:"masterItemStatus"`
			OmniStatus        string   `json:"omniStatus"`
			SalesOrganisation string   `json:"salesOrganisation"`
			SpecialAttributes []string `json:"specialAttributes"`
		} `json:"dataLayer"`
		Description struct {
			ExtraContent struct {
				Pc any `json:"pc"`
				Sp any `json:"sp"`
			} `json:"extraContent"`
			Messages struct {
				Breads   []string `json:"breads"`
				MainText string   `json:"mainText"`
				SubText  string   `json:"subText"`
				Title    string   `json:"title"`
			} `json:"messages"`
		} `json:"description"`
		Image struct {
			Details []struct {
				BackgroundColor string `json:"backgroundColor"`
				Caption         any    `json:"caption"`
				ImageURL        struct {
					Large  string `json:"large"`
					Medium string `json:"medium"`
					Small  string `json:"small"`
				} `json:"imageUrl"`
			} `json:"details"`
			Videos []struct {
				MovieURL      string `json:"movieUrl"`
				ThumbnailUrls struct {
					Large  string `json:"large"`
					Medium string `json:"medium"`
					Small  string `json:"small"`
				} `json:"thumbnailUrls"`
			} `json:"videos"`
		} `json:"image"`
		IncidentalInfo struct {
			SizeChartNote any `json:"sizeChartNote"`
		} `json:"incidentalInfo"`
		ModelCode    string `json:"modelCode"`
		Name         string `json:"name"`
		Personalizes any    `json:"personalizes"`
		Presale      any    `json:"presale"`
		Price        struct {
			Badge struct {
				Color any `json:"color"`
				Text  any `json:"text"`
			} `json:"badge"`
			Current struct {
				WithTax    string `json:"withTax"`
				WithoutTax int    `json:"withoutTax"`
			} `json:"current"`
			DiscountType string `json:"discountType"`
			Dual         any    `json:"dual"`
			Label        string `json:"label"`
		} `json:"price"`
		PurchaseInfo struct {
			Delivery struct {
				Alert        string `json:"alert"`
				ShortestDate struct {
					Day   int    `json:"day"`
					Month int    `json:"month"`
					Week  string `json:"week"`
					Year  int    `json:"year"`
				} `json:"shortestDate"`
			} `json:"delivery"`
			Messages []struct {
				IsNotice bool   `json:"isNotice"`
				Link     any    `json:"link"`
				LinkType any    `json:"linkType"`
				Text     string `json:"text"`
				Type     string `json:"type"`
			} `json:"messages"`
			News    []any `json:"news"`
			Release struct {
				Date struct {
					Day   int    `json:"day"`
					Hour  string `json:"hour"`
					Month int    `json:"month"`
					Week  string `json:"week"`
					Year  int    `json:"year"`
				} `json:"date"`
				Title string `json:"title"`
			} `json:"release"`
			StoreStockCount int `json:"storeStockCount"`
		} `json:"purchaseInfo"`
		Quantity struct {
			LimitToCart int `json:"limitToCart"`
		} `json:"quantity"`
		Skus []struct {
			AddToCartButton struct {
				Enabled bool   `json:"enabled"`
				Label   string `json:"label"`
			} `json:"addToCartButton"`
			ArticleCode  string `json:"articleCode"`
			PurchaseInfo struct {
				Icon         string `json:"icon"`
				StockMessage string `json:"stockMessage"`
			} `json:"purchaseInfo"`
			Restock struct {
				Link any `json:"link"`
			} `json:"restock"`
			SizeIndex string `json:"sizeIndex"`
			SizeName  string `json:"sizeName"`
			SkuCode   string `json:"skuCode"`
			Status    struct {
				InStockEc    bool `json:"inStockEc"`
				InStockStore bool `json:"inStockStore"`
				IsSoldOut    bool `json:"isSoldOut"`
			} `json:"status"`
		} `json:"skus"`
		Status struct {
			HasSizeTable   bool `json:"hasSizeTable"`
			HasStyling     bool `json:"hasStyling"`
			InStockEc      bool `json:"inStockEc"`
			InStockStore   bool `json:"inStockStore"`
			IsAppGuide     bool `json:"isAppGuide"`
			IsComingSoon   bool `json:"isComingSoon"`
			IsConfirmed    bool `json:"isConfirmed"`
			IsCustomize    bool `json:"isCustomize"`
			IsExactFlag    bool `json:"isExactFlag"`
			IsHype         bool `json:"isHype"`
			IsLimited      bool `json:"isLimited"`
			IsNew          bool `json:"isNew"`
			IsOutlet       bool `json:"isOutlet"`
			IsOverShipFree bool `json:"isOverShipFree"`
			IsPersonalize  bool `json:"isPersonalize"`
			IsPreorder     bool `json:"isPreorder"`
			IsPresale      bool `json:"isPresale"`
			IsRestore      bool `json:"isRestore"`
			IsReturnable   bool `json:"isReturnable"`
			IsSale         bool `json:"isSale"`
			IsSoldOut      bool `json:"isSoldOut"`
		} `json:"status"`
	} `json:"article"`
	Family any `json:"family"`
	Model  struct {
		Attributes struct {
			Age           any `json:"age"`
			BestFor       any `json:"bestFor"`
			BottomsLength any `json:"bottomsLength"`
			BraSupport    any `json:"braSupport"`
			Brand         []struct {
				Name string `json:"name"`
				Slug string `json:"slug"`
			} `json:"brand"`
			Category []struct {
				Name string `json:"name"`
				Slug string `json:"slug"`
			} `json:"category"`
			Closure []struct {
				Name string `json:"name"`
				Slug string `json:"slug"`
			} `json:"closure"`
			Collection []struct {
				Name string `json:"name"`
				Slug string `json:"slug"`
			} `json:"collection"`
			Fit []struct {
				Name string `json:"name"`
				Slug string `json:"slug"`
			} `json:"fit"`
			Gender []struct {
				Name string `json:"name"`
				Slug string `json:"slug"`
			} `json:"gender"`
			Group []struct {
				Name string `json:"name"`
				Slug string `json:"slug"`
			} `json:"group"`
			Kids                any `json:"kids"`
			ManufacturingMethod any `json:"manufacturingMethod"`
			Material            []struct {
				Name string `json:"name"`
				Slug string `json:"slug"`
			} `json:"material"`
			OtherFunction any `json:"otherFunction"`
			ShoesCut      []struct {
				Name string `json:"name"`
				Slug string `json:"slug"`
			} `json:"shoesCut"`
			SleeveLength any `json:"sleeveLength"`
			Sport        []struct {
				Name string `json:"name"`
				Slug string `json:"slug"`
			} `json:"sport"`
			Subcollection any `json:"subcollection"`
			Surface       any `json:"surface"`
			Sustainable   any `json:"sustainable"`
			Team          any `json:"team"`
			Technology    any `json:"technology"`
			Type          any `json:"type"`
		} `json:"attributes"`
		Description struct {
			ExtraContent struct {
				Pc any `json:"pc"`
				Sp any `json:"sp"`
			} `json:"extraContent"`
			Technology any `json:"technology"`
		} `json:"description"`
		ModelCode string `json:"modelCode"`
		Review    struct {
			FitbarScore int `json:"fitbarScore"`
			RatingAvg   int `json:"ratingAvg"`
			ReviewCount int `json:"reviewCount"`
			ReviewSeoLd []struct {
				Type   string `json:"@type"`
				Author struct {
					Type string `json:"@type"`
					Name string `json:"name"`
				} `json:"author"`
				DatePublished string `json:"datePublished"`
				Identifier    string `json:"identifier"`
				Name          string `json:"name"`
				ReviewBody    string `json:"reviewBody"`
				ReviewRating  struct {
					Type        string `json:"@type"`
					BestRating  int    `json:"bestRating"`
					RatingValue string `json:"ratingValue"`
				} `json:"reviewRating"`
			} `json:"reviewSeoLd"`
		} `json:"review"`
	} `json:"model"`
	RelatedArticles []struct {
		Code        string `json:"code"`
		Color       string `json:"color"`
		Image       string `json:"image"`
		IsDisplayed bool   `json:"isDisplayed"`
		Name        string `json:"name"`
		Price       struct {
			Current struct {
				WithTax    string `json:"withTax"`
				WithoutTax int    `json:"withoutTax"`
			} `json:"current"`
			DiscountType string `json:"discountType"`
			Dual         any    `json:"dual"`
			Label        string `json:"label"`
		} `json:"price"`
	} `json:"relatedArticles"`
}
