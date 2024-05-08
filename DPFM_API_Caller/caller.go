package dpfm_api_caller

import (
	"context"
	dpfm_api_input_reader "data-platform-api-sms-auth-token-generates-rmq-kube/DPFM_API_Input_Reader"
	dpfm_api_output_formatter "data-platform-api-sms-auth-token-generates-rmq-kube/DPFM_API_Output_Formatter"
	"data-platform-api-sms-auth-token-generates-rmq-kube/config"
	"github.com/latonaio/golang-logging-library-for-data-platform/logger"
	rabbitmq "github.com/latonaio/rabbitmq-golang-client-for-data-platform"
	"golang.org/x/xerrors"
	"sync"
)

type DPFMAPICaller struct {
	ctx  context.Context
	conf *config.Conf
	rmq  *rabbitmq.RabbitmqClient

	//configure    *existence_conf.ExistenceConf
	//complementer *sub_func_complementer.SubFuncComplementer
}

func NewDPFMAPICaller(
	conf *config.Conf,
	rmq *rabbitmq.RabbitmqClient,
	// confirmor *existence_conf.ExistenceConf,
	// complementer *sub_func_complementer.SubFuncComplementer,
) *DPFMAPICaller {
	return &DPFMAPICaller{
		ctx:  context.Background(),
		conf: conf,
		rmq:  rmq,
		//configure:    confirmor,
		//complementer: complementer,
	}
}

func (c *DPFMAPICaller) AsyncCreates(
	accepter []string,
	input *dpfm_api_input_reader.SDC,
	output *dpfm_api_output_formatter.SDC,
	log *logger.Logger,
) (interface{}, []error) {
	var smsAuthToken *[]dpfm_api_output_formatter.SMSAuthToken
	var errs []error

	for _, fn := range accepter {
		switch fn {
		case "SMSAuthToken":
			smsAuthToken = c.SMSAuthToken(input, &errs, log, c.conf)
		default:
		}
	}

	data := &dpfm_api_output_formatter.Message{
		SMSAuthToken: smsAuthToken,
	}

	return data, nil
}

func (c *DPFMAPICaller) exconfProcess(
	mtx *sync.Mutex,
	wg *sync.WaitGroup,
	exconfFin chan error,
	input *dpfm_api_input_reader.SDC,
	output *dpfm_api_output_formatter.SDC,
	exconfAllExist *bool,
	accepter []string,
	errs *[]error,
	log *logger.Logger,
) {
	defer wg.Done()
	var e []error
	//*exconfAllExist, e = c.configure.Conf(input, output, accepter, log)
	if len(e) != 0 {
		mtx.Lock()
		*errs = append(*errs, e...)
		mtx.Unlock()
		exconfFin <- xerrors.New("exconf error")
		return
	}
	exconfFin <- nil
}
