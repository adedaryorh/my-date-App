// coverage:ignore-file
// GENERATED CODE - DO NOT MODIFY BY HAND
// ignore_for_file: type=lint
// ignore_for_file: unused_element, deprecated_member_use, deprecated_member_use_from_same_package, use_function_type_syntax_for_parameters, unnecessary_const, avoid_init_to_null, invalid_override_different_default_values_named, prefer_expression_function_bodies, annotate_overrides, invalid_annotation_target, unnecessary_question_mark

part of '../post_model.dart';

// **************************************************************************
// FreezedGenerator
// **************************************************************************

T _$identity<T>(T value) => value;

final _privateConstructorUsedError = UnsupportedError(
    'It seems like you constructed your class using `MyClass._()`. This constructor is only meant to be used by freezed and you are not supposed to need it nor use it.\nPlease check the documentation here for more information: https://github.com/rrousselGit/freezed#adding-getters-and-methods-to-our-models');

PostModel _$PostModelFromJson(Map<String, dynamic> json) {
  return _PostModel.fromJson(json);
}

/// @nodoc
mixin _$PostModel {
  String? get userName => throw _privateConstructorUsedError;
  String? get time => throw _privateConstructorUsedError;
  String? get content => throw _privateConstructorUsedError;
  List<String> get media => throw _privateConstructorUsedError;
  bool get isVideo => throw _privateConstructorUsedError;

  Map<String, dynamic> toJson() => throw _privateConstructorUsedError;
  @JsonKey(ignore: true)
  $PostModelCopyWith<PostModel> get copyWith =>
      throw _privateConstructorUsedError;
}

/// @nodoc
abstract class $PostModelCopyWith<$Res> {
  factory $PostModelCopyWith(PostModel value, $Res Function(PostModel) then) =
      _$PostModelCopyWithImpl<$Res, PostModel>;
  @useResult
  $Res call(
      {String? userName,
      String? time,
      String? content,
      List<String> media,
      bool isVideo});
}

/// @nodoc
class _$PostModelCopyWithImpl<$Res, $Val extends PostModel>
    implements $PostModelCopyWith<$Res> {
  _$PostModelCopyWithImpl(this._value, this._then);

  // ignore: unused_field
  final $Val _value;
  // ignore: unused_field
  final $Res Function($Val) _then;

  @pragma('vm:prefer-inline')
  @override
  $Res call({
    Object? userName = freezed,
    Object? time = freezed,
    Object? content = freezed,
    Object? media = null,
    Object? isVideo = null,
  }) {
    return _then(_value.copyWith(
      userName: freezed == userName
          ? _value.userName
          : userName // ignore: cast_nullable_to_non_nullable
              as String?,
      time: freezed == time
          ? _value.time
          : time // ignore: cast_nullable_to_non_nullable
              as String?,
      content: freezed == content
          ? _value.content
          : content // ignore: cast_nullable_to_non_nullable
              as String?,
      media: null == media
          ? _value.media
          : media // ignore: cast_nullable_to_non_nullable
              as List<String>,
      isVideo: null == isVideo
          ? _value.isVideo
          : isVideo // ignore: cast_nullable_to_non_nullable
              as bool,
    ) as $Val);
  }
}

/// @nodoc
abstract class _$$PostModelImplCopyWith<$Res>
    implements $PostModelCopyWith<$Res> {
  factory _$$PostModelImplCopyWith(
          _$PostModelImpl value, $Res Function(_$PostModelImpl) then) =
      __$$PostModelImplCopyWithImpl<$Res>;
  @override
  @useResult
  $Res call(
      {String? userName,
      String? time,
      String? content,
      List<String> media,
      bool isVideo});
}

/// @nodoc
class __$$PostModelImplCopyWithImpl<$Res>
    extends _$PostModelCopyWithImpl<$Res, _$PostModelImpl>
    implements _$$PostModelImplCopyWith<$Res> {
  __$$PostModelImplCopyWithImpl(
      _$PostModelImpl _value, $Res Function(_$PostModelImpl) _then)
      : super(_value, _then);

  @pragma('vm:prefer-inline')
  @override
  $Res call({
    Object? userName = freezed,
    Object? time = freezed,
    Object? content = freezed,
    Object? media = null,
    Object? isVideo = null,
  }) {
    return _then(_$PostModelImpl(
      userName: freezed == userName
          ? _value.userName
          : userName // ignore: cast_nullable_to_non_nullable
              as String?,
      time: freezed == time
          ? _value.time
          : time // ignore: cast_nullable_to_non_nullable
              as String?,
      content: freezed == content
          ? _value.content
          : content // ignore: cast_nullable_to_non_nullable
              as String?,
      media: null == media
          ? _value._media
          : media // ignore: cast_nullable_to_non_nullable
              as List<String>,
      isVideo: null == isVideo
          ? _value.isVideo
          : isVideo // ignore: cast_nullable_to_non_nullable
              as bool,
    ));
  }
}

/// @nodoc
@JsonSerializable()
class _$PostModelImpl implements _PostModel {
  const _$PostModelImpl(
      {this.userName = '',
      this.time = '',
      this.content = '',
      final List<String> media = const [],
      this.isVideo = false})
      : _media = media;

  factory _$PostModelImpl.fromJson(Map<String, dynamic> json) =>
      _$$PostModelImplFromJson(json);

  @override
  @JsonKey()
  final String? userName;
  @override
  @JsonKey()
  final String? time;
  @override
  @JsonKey()
  final String? content;
  final List<String> _media;
  @override
  @JsonKey()
  List<String> get media {
    if (_media is EqualUnmodifiableListView) return _media;
    // ignore: implicit_dynamic_type
    return EqualUnmodifiableListView(_media);
  }

  @override
  @JsonKey()
  final bool isVideo;

  @override
  String toString() {
    return 'PostModel(userName: $userName, time: $time, content: $content, media: $media, isVideo: $isVideo)';
  }

  @override
  bool operator ==(Object other) {
    return identical(this, other) ||
        (other.runtimeType == runtimeType &&
            other is _$PostModelImpl &&
            (identical(other.userName, userName) ||
                other.userName == userName) &&
            (identical(other.time, time) || other.time == time) &&
            (identical(other.content, content) || other.content == content) &&
            const DeepCollectionEquality().equals(other._media, _media) &&
            (identical(other.isVideo, isVideo) || other.isVideo == isVideo));
  }

  @JsonKey(ignore: true)
  @override
  int get hashCode => Object.hash(runtimeType, userName, time, content,
      const DeepCollectionEquality().hash(_media), isVideo);

  @JsonKey(ignore: true)
  @override
  @pragma('vm:prefer-inline')
  _$$PostModelImplCopyWith<_$PostModelImpl> get copyWith =>
      __$$PostModelImplCopyWithImpl<_$PostModelImpl>(this, _$identity);

  @override
  Map<String, dynamic> toJson() {
    return _$$PostModelImplToJson(
      this,
    );
  }
}

abstract class _PostModel implements PostModel {
  const factory _PostModel(
      {final String? userName,
      final String? time,
      final String? content,
      final List<String> media,
      final bool isVideo}) = _$PostModelImpl;

  factory _PostModel.fromJson(Map<String, dynamic> json) =
      _$PostModelImpl.fromJson;

  @override
  String? get userName;
  @override
  String? get time;
  @override
  String? get content;
  @override
  List<String> get media;
  @override
  bool get isVideo;
  @override
  @JsonKey(ignore: true)
  _$$PostModelImplCopyWith<_$PostModelImpl> get copyWith =>
      throw _privateConstructorUsedError;
}
