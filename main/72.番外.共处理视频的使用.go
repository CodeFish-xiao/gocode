package main

import (
	"github.com/nareix/joy4/av"
	"github.com/nareix/joy4/av/avutil"
	"github.com/nareix/joy4/av/transcode"
	"github.com/nareix/joy4/cgo/ffmpeg"
	"github.com/nareix/joy4/format"
)

// need ffmpeg with libfdkaac installed

func init() {
	format.RegisterAll()
}

func main() {
	infile, _ := avutil.Open("speex.flv")

	findcodec := func(stream av.AudioCodecData, i int) (need bool, dec av.AudioDecoder, enc av.AudioEncoder, err error) {
		need = true
		dec, _ = ffmpeg.NewAudioDecoder(stream)
		enc, _ = ffmpeg.NewAudioEncoderByName("libfdk_aac")
		_ = enc.SetSampleRate(stream.SampleRate())
		_ = enc.SetChannelLayout(av.CH_STEREO)
		_ = enc.SetBitrate(12000)
		_ = enc.SetOption("profile", "HE-AACv2")
		return
	}

	trans := &transcode.Demuxer{
		Options: transcode.Options{
			FindAudioDecoderEncoder: findcodec,
		},
		Demuxer: infile,
	}

	outfile, _ := avutil.Create("out.ts")
	_ = avutil.CopyFile(outfile, trans)

	_ = outfile.Close()
	_ = infile.Close()
	_ = trans.Close()
}
