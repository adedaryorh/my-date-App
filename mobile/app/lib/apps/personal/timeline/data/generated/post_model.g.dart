// GENERATED CODE - DO NOT MODIFY BY HAND

part of '../post_model.dart';

// **************************************************************************
// JsonSerializableGenerator
// **************************************************************************

_$PostModelImpl _$$PostModelImplFromJson(Map<String, dynamic> json) =>
    _$PostModelImpl(
      userName: json['userName'] as String? ?? '',
      time: json['time'] as String? ?? '',
      content: json['content'] as String? ?? '',
      media:
          (json['media'] as List<dynamic>?)?.map((e) => e as String).toList() ??
              const [],
      isVideo: json['isVideo'] as bool? ?? false,
    );

Map<String, dynamic> _$$PostModelImplToJson(_$PostModelImpl instance) =>
    <String, dynamic>{
      'userName': instance.userName,
      'time': instance.time,
      'content': instance.content,
      'media': instance.media,
      'isVideo': instance.isVideo,
    };
