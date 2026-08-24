package report

import (
	"encoding/csv"
	"fmt"
	"github.com/jb843051627/mangrove-flux/internal/model"
	"io"
	"strconv"
)

func WriteCSV(w io.Writer, reports []model.FluxReport) error {
	c := csv.NewWriter(w)
	if err := c.Write([]string{"station_id", "day", "samples", "good_samples", "mean_co2", "mean_ch4", "net_carbon", "alerts"}); err != nil {
		return err
	}
	for _, item := range reports {
		row := []string{item.StationID, item.Day, strconv.Itoa(item.Samples), strconv.Itoa(item.GoodSamples), fmt.Sprintf("%.4f", item.MeanCO2), fmt.Sprintf("%.4f", item.MeanCH4), fmt.Sprintf("%.4f", item.NetCarbon), strconv.Itoa(item.AlertCount)}
		if err := c.Write(row); err != nil {
			return err
		}
	}
	c.Flush()
	return c.Error()
}
